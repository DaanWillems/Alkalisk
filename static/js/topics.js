var app = new Vue({
    el: '#forum',
    data: function() {
        return{
            topics: []
        }
    },
    created: function () {
        this.$http.get('/getTopics').then(function(response) {
        console.log(response.data);
            this.topics = response.data ? response.data : []
        })
    },
})