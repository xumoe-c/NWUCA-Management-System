package main

import (
	"NWUCA-Management-System/server/config"
	"NWUCA-Management-System/server/internal/handler"
	"NWUCA-Management-System/server/internal/middleware"
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
	"NWUCA-Management-System/server/internal/service"
	"log"

	_ "NWUCA-Management-System/server/docs" // 导入 docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title NWUCA Management System API
// @version 1.0
// @description This is the API documentation for the NWUCA Management System.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 1. 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// 2. 初始化数据库
	db, err := config.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("could not init database: %v", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Member{}, &model.Department{}, &model.Assignment{}, &model.Advisor{}, &model.Position{})

	// 3. 依赖注入：从内到外创建实例
	// Repository -> Service -> Handler
	userRepo := repository.NewUserRepository(db)
	departmentRepo := repository.NewDepartmentRepository(db)
	positionRepo := repository.NewPositionRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	assignmentRepo := repository.NewAssignmentRepository(db)

	// 依赖注入：将仓库注入服务
	userService := service.NewUserService(userRepo, cfg.JWT.SecretKey, cfg.JWT.ExpirationDays)
	departmentService := service.NewDepartmentService(departmentRepo)
	positionService := service.NewPositionService(positionRepo)
	memberService := service.NewMemberService(db, memberRepo, userRepo)
	assignmentService := service.NewAssignmentService(assignmentRepo)

	// 依赖注入：将服务注入处理器
	userHandler := handler.NewUserHandler(userService)
	departmentHandler := handler.NewDepartmentHandler(departmentService)
	positionHandler := handler.NewPositionHandler(positionService)
	memberHandler := handler.NewMemberHandler(memberService)
	assignmentHandler := handler.NewAssignmentHandler(assignmentService)

	// 4. 初始化 Gin 引擎
	r := gin.Default()

	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 5. 注册路由
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/register", userHandler.Register)
		apiV1.POST("/login", userHandler.Login)

		// 受保护的路由组
		protected := apiV1.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWT.SecretKey))
		{
			// 部门管理 - 仅限管理员
			adminDepartments := protected.Group("/departments")
			adminDepartments.Use(middleware.RoleAuthMiddleware("admin"))
			{
				adminDepartments.POST("/", departmentHandler.CreateDepartment)
				adminDepartments.PUT("/:id", departmentHandler.UpdateDepartment)
				adminDepartments.DELETE("/:id", departmentHandler.DeleteDepartment)
			}
			// 部门查看 - 所有认证用户
			protected.GET("/departments", departmentHandler.GetDepartments)

			// 职务管理 - 仅限管理员
			adminPositions := protected.Group("/positions")
			adminPositions.Use(middleware.RoleAuthMiddleware("admin"))
			{
				adminPositions.POST("/", positionHandler.CreatePosition)
				adminPositions.PUT("/:id", positionHandler.UpdatePosition)
				adminPositions.DELETE("/:id", positionHandler.DeletePosition)
			}
			// 职务查看 - 所有认证用户
			protected.GET("/positions", positionHandler.GetPositions)

			// 会员管理 - 仅限管理员
			adminMembers := protected.Group("/members")
			adminMembers.Use(middleware.RoleAuthMiddleware("admin"))
			{
				adminMembers.POST("/", memberHandler.CreateMember)
				adminMembers.PUT("/:id", memberHandler.UpdateMember)
				adminMembers.DELETE("/:id", memberHandler.DeleteMember)
			}
			// 会员查看 - 所有认证用户
			protected.GET("/members", memberHandler.GetMembers)

			// 任期管理 - 仅限管理员
			adminAssignments := protected.Group("/assignments")
			adminAssignments.Use(middleware.RoleAuthMiddleware("admin"))
			{
				adminAssignments.POST("/", assignmentHandler.CreateAssignment)
				adminAssignments.PUT("/:id", assignmentHandler.UpdateAssignment)
				adminAssignments.DELETE("/:id", assignmentHandler.DeleteAssignment)
			}
			// 任期查看 - 所有认证用户
			protected.GET("/assignments", assignmentHandler.GetAssignments)
		}
	}

	// 6. 启动服务
	if cfg.Server.Port == "" {
		cfg.Server.Port = ":8080" // 默认端口
	}
	err = r.Run(cfg.Server.Port)
}
