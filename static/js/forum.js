var Forum = Vue.component("forum", {
    template: `
     <div id="forum">
     <router-link :to="'/createPost'" ><a class="btn-floating btn-large waves-effect waves-light red"><i class="material-icons">add</i></a></router-link>
        <div id="posts">
            <div class="topic" v-for="post in posts">
                <div class="thumbnail">
                    <div class="title">{{ post.Title }}</div>
                </div>
                <i class="last-posted material-icons">access_time</i><p>{{ post.LastCommentAt }}</p>
                 <router-link :to="'/post/' + post.Id" ><i class="material-icons align-right">forum</i></router-link>
            </div>
        </div>
    </div>
    `,
    data: function() {
        return{
            posts: []
        }
    },
    created: function () {
        var self = this;
        $.get("/getPosts", function(data) {
            self.posts = JSON.parse(data);
        })
    } 
})
