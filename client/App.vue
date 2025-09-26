<script>
  import { isLoggedIn } from './store/auth'
	export default {
		onLaunch: function() {
			console.log('App Launch')
			// Redirect to login if not logged in
			const pages = getCurrentPages();
			const current = pages && pages.length ? pages[pages.length - 1] : null;
			const route = current && (current.route || (current.$page && current.$page.path));
			if (!isLoggedIn() && route !== 'pages/login/index') {
				uni.reLaunch({ url: '/pages/login/index' })
			}
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style>
	/*每个页面公共css */
</style>
