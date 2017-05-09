var Post = Vue.component("forum", {
    template: `
    <div>
        <div id="post">
            <router-link :to="'/'"><a class="btn-floating btn-large waves-effect waves-light back-button red"><i class="material-icons">keyboard_backspace</i></a></router-link>
            <i class="menu material-icons">more_vert</i>
            <div class="header">
                <img class="profile-pic" src="/img/anon.png">
                <div id="title">{{post.Title}}</div><br>
            </div>
            <div class="content">{{post.Content}}</div>
        </div>
        <div id="comments"> 
            <div id="new-comment">
                <div class="input-field col s12">
                    <textarea id="new-comment-text" class="materialize-textarea"></textarea>
                    <label for="new-comment-text">Type your comment here</label>
                </div>
                <div v-on:click="postComment" class="waves-effect waves-light btn">Comment</div>
            </div>
            <div v-for="comment in comments" class="comment">
                <div class="header">
                    <img class="profile-pic" src="/img/anon.png">
                    <div id="title">{{comment.CreatedAt}}</div><br>
                </div>
                <div class="content">{{comment.Content}}</div><br>
            </div>
        </div>
    </div>
    `,
    data: function() {
        return{
            post: {},
            comments: []
        }
    },
    created: function () {
        var self = this;
        var id = this.$route.params.id;
        $.get("/getPost/"+id, function(data) {
            self.post = JSON.parse(data);
        })
        $.get("/getComments/"+id, function(data) {
            self.comments = JSON.parse(data);
        })
    },
    methods: {
        postComment: function() {
            var self = this;
            $.ajax({
                type: "POST",
                url: "/postComment/"+this.$route.params.id,
                data: JSON.stringify({
                    content: $("#new-comment-text").val(),
                    id: self.$route.params.id
                }),
                success: function(data) {
                    if(self.comments == null) {
                        self.comments = [data]
                        $("#new-comment-text").val("");
                     } else {
                        self.comments.unshift(data);
                        $("#new-comment-text").val("");
                     }
                },
                contentType: "application/json; charset=utf-8",
                dataType: "json"
            });
        } 
    } 
})
