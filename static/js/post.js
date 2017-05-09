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
            <div id="content">{{post.Content}}</div>
        </div>
        <div id="comments"> 
            <div v-for="comment in comments" class="comment">
                <div class="header">
                    <img class="profile-pic" src="/img/anon.png">
                    <div id="title">{{comment.Title}}</div><br>
                </div>
            </div>
        </div>
    </div>
    `,
    data: function() {
        return{
            post: {},
            comments: [{"Title": "test"}, {"Title": "test1"}]
        }
    },
    created: function () {
        var self = this;
        var id = this.$route.params.id;
        $.get("/getPost/"+id, function(data) {
            self.post = JSON.parse(data);
        })
    } 
})
