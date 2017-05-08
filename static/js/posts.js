const Post = { 
    template: `
    <div id="posts">
      <div class="topic" v-for="post in posts">
            <div class="thumbnail">
                <div class="title">{{ post.Title }}</div>
            </div>
            <i class="last-posted material-icons">access_time</i><p>{{ post.LastCommentAt }}</p>
            <i class="material-icons align-right">forum</i>
        </div>
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