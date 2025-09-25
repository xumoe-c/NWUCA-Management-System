# Epic1/2 前端页面设计

本文档规划 Epic 1（核心系统与用户认证）与 Epic 2（组织与人员管理）的 H5 Web 端前端页面原型与交互规范，统一使用 TDesign 作为视觉设计系统与组件库，以便快速产出稳定的一致性体验。

> 说明：MVP 阶段优先交付 H5 Web 端（TDesign Vue Next）。后续小程序/移动端可基于 tdesign-mobile-vue 或自定义组件做适配。

---

## 1. 目标与范围

- 目标：
  - 用最短路径完成“认证 + 组织与人员管理”端到端闭环。
  - 输出可直接落地的页面结构、组件选型与交互细节，降低实现沟通成本。
- 范围：
  - Epic 1：登录/登出、路由守卫、基础布局与导航、个人中心概览。
  - Epic 2：部门/职位/成员/任期的列表、创建/编辑、详情预览与批量操作（可部分延后）。

---

## 2. 设计系统与主题

- 组件库：TDesign Vue Next（Vue3）。
- 主题与品牌：
  - 品牌色建议：`--td-brand-color: #2A82E4`（可根据社团视觉规范调整）。
  - 暗黑模式：非必需；优先亮色主题，预留开关入口在头像下拉菜单。
- 全局交互：
  - 反馈：`t-message`（轻提示），`t-notification`（重要提醒），`t-dialog`（危险操作二次确认）。
  - 表单：`t-form` + 校验规则；提交中显示 `submitLoading`，禁用二次点击。
  - 列表：`t-table` + `t-pagination`；支持服务端分页、排序、过滤。
  - 日期：`t-date-picker`；时间范围用 `mode="date-range"`。
  - 上传：`t-upload`（头像/附件）；统一文件大小/类型限制与错误提示。

---

## 3. 信息架构（IA）与路由

- 顶层导航（侧边栏）：

  - 仪表盘（Dashboard）
  - 组织管理
    - 部门管理
    - 职务管理
    - 成员管理
    - 任期分配
  - 个人中心
- 顶部栏：面包屑、搜索（可选）、全局通知、用户头像/菜单（登出、主题、个人中心）。
- 路由建议（uni-app H5，通过 `pages.json` 配置；此处为示意路径）：

  - `/pages/login/index` 登录
  - `/pages/dashboard/index` 仪表盘
  - `/pages/departments/index` 部门管理
  - `/pages/positions/index` 职务管理
  - `/pages/members/index` 成员管理
  - `/pages/assignments/index` 任期分配
  - `/pages/profile/index` 个人中心
- 路由守卫：

  - 未登录：重定向至 `/pages/login/index`。
  - 已登录但无权限：显示 403 结果页（`t-result`），提供返回首页按钮。

---

## 4. 公共布局与基础组件

- 布局骨架：
  - `t-layout`（Header / Sider / Content）或自定义 Grid：
    - Header：Logo + 系统名、面包屑、搜索（可选）、通知、用户菜单。
    - Sider：`t-menu` + 图标（使用 TDesign Icons）。
    - Content：`t-breadcrumb` + 页面标题 + 操作区（按钮组）+ 主体。
- 通用组件：
  - 页面标题区：标题、副标题、右上角“新增”等操作按钮。
  - 查询区（可折叠）：`t-form` + 行内布局；提供“重置/查询”按钮。
  - 主表格：`t-table`；建议开启列设置（显示/隐藏）、自适应宽度、空态 `t-empty`。
  - 侧滑抽屉：`t-drawer` 用于详情与编辑；保持上下文不丢失。
  - 确认对话框：`t-dialog` 二次确认删除/危险操作。

---

## 5. 页面原型（Epic 1）

### 5.1 登录页（/pages/login/index）

- 目标：账号/密码登录，获取 JWT 并保存（Pinia/本地存储）。
- 结构：居中卡片 `t-card` 内为 `t-form`。
- 表单字段：
  - 账号（`username`）：`t-input`，必填。
  - 密码（`password`）：`t-input type="password"`，必填。
  - 记住我（可选）：`t-checkbox`（仅 H5）。
- 操作：
  - 登录：`t-button theme="primary" block`，提交校验，loading 态。
- 反馈：
  - 成功：跳转 Dashboard，并 `t-message.success("登录成功")`。
  - 失败：解析错误码，`t-message.error("用户名或密码错误")`。
- 状态与异常：
  - 登录中禁用按钮与输入；401 全局拦截提示后重定向登录。
- 接口：`POST /api/v1/login`（参考后端登录接口；若尚未定义，需统一）。

### 5.2 仪表盘（/pages/dashboard/index）

- 目标：快速入口与概览数据。
- 模块：
  - 快捷入口卡片（成员、部门、职位、任期）`t-card-grid`。
  - 待办（如后续引入审批/通知）占位 `t-list`。
  - 最近操作/最近创建成员 `t-table`（简版）。
- 说明：MVP 可先放置占位组件 + 引导文案，后续迭代接入真实数据。

### 5.3 个人中心（/pages/profile/index）

- 目标：展示当前登录用户信息，支持密码修改（跳转二级页或弹窗）。
- 模块：
  - 用户信息卡：头像（`t-avatar` + `t-upload`）、姓名、角色、邮箱。
  - 安全设置：修改密码按钮，弹出 `t-dialog` + `t-form`（原密码/新密码/确认新密码）。
- 接口：
  - `GET /api/v1/users/me`（拉取个人信息）。
  - `PUT /api/v1/users/me/password`（修改密码）。

---

## 6. 页面原型（Epic 2）

### 6.1 部门管理（/pages/departments/index）

- 目标：增删改查部门，支持树结构与排序（排序可延后）。
- 布局：
  - 左侧：部门树 `t-tree`（支持搜索、展开/收起、选中节点）。
  - 右侧：选中部门详情卡 + 子部门列表 `t-table`。
- 交互：
  - 新增部门：顶部“新增部门”按钮 `t-button` -> `t-drawer` 表单。
  - 编辑部门：在详情卡或子部门表格的“编辑”按钮 -> `t-drawer`。
  - 删除：`t-dialog` 确认（需校验是否存在子部门/成员）。
- 表单字段：
  - 名称 `name`（必填）、描述 `description`、上级部门 `parent_id`（`t-tree-select`）。
- 权限：
  - 浏览：登录用户可见。
  - 新增/编辑/删除：仅 `admin` 显示操作按钮。
- 接口：
  - `GET /api/v1/departments` 列表/树。
  - `POST/PUT/DELETE /api/v1/departments` 变更。

### 6.2 职务管理（/pages/positions/index）

- 目标：维护职位信息。
- 结构：
  - 查询区：关键字（名称）、访问级别区间（可选）。
  - 列表：`t-table`（名称、访问级别、创建/更新时间、操作）。
- 交互：
  - 新增/编辑：`t-drawer` 表单。
  - 删除：`t-dialog` 确认。
- 表单字段：
  - 名称 `name`（必填）、访问级别 `access_level`（数字，必填）、描述（可选）。
- 权限：
  - 浏览：登录用户可见。
  - 新增/编辑/删除：仅 `admin`。
- 接口：
  - `GET /api/v1/positions`
  - `POST/PUT/DELETE /api/v1/positions`

### 6.3 成员管理（/pages/members/index）

- 目标：成员档案的列表、创建、编辑、导出（导出可延后）。
- 结构：
  - 查询区：
    - 关键字（姓名/学号/电话/邮箱）
    - 成员状态（在读/已毕业/休学/退会）`t-select`
    - 年级 `grade`、学院 `college`、专业 `major`
    - 部门 `department`（`t-tree-select`，基于最新任期）
  - 列表：`t-table`
    - 列：头像、姓名、学号、年级、学院/专业、状态、加入时间、当前部门/职务、操作。
  - 操作区：新增成员、批量导入（后续）、导出（后续）。
- 交互：
  - 查看详情：行点击或“查看” -> `t-drawer` 展示详情 + 任期时间线 `t-timeline`。
  - 新增/编辑：`t-drawer` 表单（支持创建关联 `User`）。
  - 删除：`t-dialog` 确认（若存在在任任期，需阻止并提示）。
- 表单字段（核心）：
  - 姓名、学号、年级、学院、专业、电话、邮箱、入会时间、预计毕业时间、状态、头像上传、技能标签（后续）、备注。
- 权限：
  - 浏览：登录用户可见。
  - 新增/编辑/删除：仅 `admin`。
- 接口：
  - `GET /api/v1/members`
  - `POST/PUT/DELETE /api/v1/members`

### 6.4 任期分配（/pages/assignments/index）

- 目标：为成员在特定部门/队伍分配职务与起止时间；查看历史记录。
- 结构：
  - 查询区：成员、部门（树）、职务、是否在任（`EndDate` 为空）。
  - 列表：`t-table`
    - 列：成员、部门/队伍、职务、开始日期、结束日期（为空显示“在任”标记）、操作。
  - 操作区：新增任期。
- 交互：
  - 新增/编辑：`t-drawer` 表单；选择成员（`t-select` 搜索）、部门（`t-tree-select`）、职务（`t-select`）、开始/结束日期（`t-date-picker`）。
  - 删除：`t-dialog` 确认（影响权限，显示风险提示）。
- 权限：
  - 浏览：登录用户可见。
  - 新增/编辑/删除：仅 `admin`。
- 接口：
  - `GET /api/v1/assignments`
  - `POST/PUT/DELETE /api/v1/assignments`

---

## 7. 表单与校验规范

- 必填项显示星号；错误提示在项下方展示，提交时滚动到首个错误项。
- 通用校验：
  - 手机号：`/^1\d{10}$/`（可根据学校规范调整）。
  - 邮箱：标准邮箱格式校验。
  - 学号：长度与字符集校验（按校内规范）。
  - 日期：开始日期不得晚于结束日期。
- 提交行为：
  - 提交前禁用按钮，显示 loading；成功关闭抽屉并刷新列表；失败保留表单并提示。

---

## 8. 表格、筛选与状态规范

- 表格：
  - 行操作区域统一使用 `t-space` 间距；图标按钮用于次要操作。
  - 空态使用 `t-empty`，提供“新增”快捷入口。
- 筛选：
  - 查询区支持折叠/展开；支持“重置”和“查询”。
  - 长列表筛选建议服务端执行；分页尺寸可选 10/20/50。
- 状态标识：
  - 成员状态：`t-tag`（在读=primary、已毕业=success、休学=warning、退会=danger）。
  - 任期状态：结束日期为空显示 `t-tag`（“在任”）。

---

## 9. 权限与可见性

- 角色：`admin` 与 `member`（MVP）。
- 可见性：
  - 列表与详情：登录可见。
  - 写操作按钮：仅 `admin` 渲染或启用；后端同时二次校验。
- 路由守卫：在 `onLaunch` 或路由拦截中检查 Token；401 清理态并跳转登录。

---

## 10. 接口映射与数据约定（与后端对齐）

- 统一约定：
  - 所有受保护接口在 `Authorization: Bearer <token>` 下访问。
  - 统一响应结构参考 `internal/dto/response_dto.go`（例如包含 `code`、`message`、`data`）。
- 主要资源：
  - Departments：`GET/POST/PUT/DELETE /api/v1/departments`
  - Positions：`GET/POST/PUT/DELETE /api/v1/positions`
  - Members：`GET/POST/PUT/DELETE /api/v1/members`
  - Assignments：`GET/POST/PUT/DELETE /api/v1/assignments`
  - Auth：`POST /api/v1/login`（如未定义，需与后端统一）

---

## 11. 加载/错误/空态与反馈

- 加载：页面与表格加载使用骨架屏或内置 loading；表单提交使用按钮 loading。
- 错误：
  - 401：清理登录态并跳转登录；提示“登录已过期”。
  - 403：展示 `t-result` 禁止访问，提供返回首页。
  - 5xx：`t-notification.error`，提示“服务繁忙，请稍后重试”。
- 空态：`t-empty` 配合行动按钮（如“新增成员”）。

---

## 12. 移动端适配（H5）

- 列表在小屏幕隐藏次要列，提供“列设置”让用户自定义显示。
- 抽屉在移动端全宽显示；对话框使用底部抽屉样式（可用 CSS 覆盖）。
- 查询区默认折叠，仅露出 1-2 个关键筛选项。

---

## 13. 实现清单（MVP Sprint）

- 公共：
  - 基础布局与导航（Header/Sider/Content），主题与消息组件接入。
  - Axios 封装：请求拦截器注入 JWT，响应拦截器处理 401/403。
  - 路由守卫：未登录跳登录；403 结果页。
- 页面：
  - 登录页、仪表盘、个人中心。
  - 部门管理：树 + 子表格 + 抽屉表单。
  - 职务管理：表格 + 抽屉表单。
  - 成员管理：查询 + 表格 + 详情抽屉 + 表单。
  - 任期分配：查询 + 表格 + 表单。

---

## 14. 风险与备选方案

- uni-app 与 TDesign 兼容性：MVP 先专注 H5 Web（TDesign Vue Next），小程序端后续用 tdesign-miniprogram 或移动端组件替换。
- 权限细粒度：当前以角色级（admin/member）为主；后续升级至“权限点”后，前端按钮显隐需绑定细粒度权限。
- 数据一致性与约束：删除成员/部门时需校验关联（任期/子部门/用户），前端需基于接口返回做明确提示与拦截。

---

## 15. 附：页面简要线框（文字稿）

- 成员管理列表：
  - 顶部：标题“成员管理” + [新增成员]
  - 查询区（可折叠）：关键字 | 状态 | 年级 | 部门 | [重置] [查询]
  - 表格：头像 | 姓名 | 学号 | 年级 | 学院/专业 | 状态 | 加入时间 | 当前部门/职务 | 操作（查看/编辑/删除）
  - 抽屉（查看/编辑）：基础信息 | 任期时间线 | 操作区（保存/取消）
- 任期分配列表：
  - 顶部：标题“任期分配” + [新增任期]
  - 查询区：成员 | 部门 | 职务 | 是否在任 | [重置] [查询]
  - 表格：成员 | 部门/队伍 | 职务 | 开始日期 | 结束日期（在任） | 操作（编辑/删除）
