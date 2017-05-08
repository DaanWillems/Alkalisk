var app = new Vue({
    el: '#forum',
    data: function() {
        return{
            posts: []
        }
    },
    created: function () {
        this.$http.get('/getPosts').then(function(response) {
        console.log(response.data);
            this.posts = response.data ? response.data : []
        })
    },
})