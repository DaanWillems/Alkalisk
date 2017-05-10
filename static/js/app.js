$(document).ready(function(){
  $('.button-collapse').sideNav();
});

var routes = [
  { path: '/', component: Forum },
  { path: '/post/:id', component: Post },
  { path: '/createPost', component: NewPost },
  { path: '*', redirect: '/'}
]

var router = new VueRouter({
  mode: 'history',
  routes // short for routes: routes
})

var app = new Vue({
  router
}).$mount('#app')