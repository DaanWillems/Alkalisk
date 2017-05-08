$(document).ready(function(){
  $('.button-collapse').sideNav();
});

const routes = [
  { path: '/comment', component: Comments },
  { path: '/', component: Post }
]

const router = new VueRouter({
  routes // short for routes: routes
})

const app = new Vue({
  router
}).$mount('#forum')