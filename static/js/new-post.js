var NewPost = Vue.component("new-post", {
    template: `
    <div id="new-post">
        <div class="input-field col s12">
            <input id="new-post-title" type="text" class="validate">
            <label for="new-post-title">Title</label>
        </div>
        <div class="input-field col s12">
            <textarea id="new-post-text" class="materialize-textarea"></textarea>
            <label for="new-post-text">Type your comment here</label>
        </div>
        <div v-on:click="savePost" class="waves-effect waves-light btn">post</div>
    </div>
    `,
    methods: {
        savePost: function() {
            var self = this;
            $.ajax({
                type: "POST",
                url: "/newPost",
                data: JSON.stringify({
                    title: $("#new-post-title").val(),
                    content: $("#new-post-text").val()
                }),
                success: function(data) {
                    console.log("tessst")
                    router.push('/')
                },
                contentType: "application/json; charset=utf-8",
                dataType: "json"
            });
            
        }
    }
})