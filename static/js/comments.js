const Comments = { 
    template: `
    <div id="comments">
        test
    </div>
    `,
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
    } 
}