import React from 'react';

class Post extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            permalink: this.props.match.params.permalink,
            post: null,
            postView: 0,
            totalView: 0,
        };
    }

    componentDidMount() {
        var api = '/api/v1/search/';
        if (window.location.hostname === 'localhost') {
            api = 'http://localhost:8080/api/v1/search/';
        }
        fetch(api + this.state.permalink)
            .then(res => res.json())
            .then((data) => {
                this.setState({
                    post: data.post,
                    postView: data.postView,
                    totalView: data.totalView,
                });

                // Enable codepen.
                const codepen = document.createElement('script');
                codepen.src = "https://static.codepen.io/assets/embed/ei.js";
                codepen.async = true;
                document.body.appendChild(codepen);

                // Enable discourse.
                const discourse = document.createElement('script');
                discourse.innerHTML =  "var DiscourseEmbed = { discourseUrl: 'https://discourse.wwei10.com/',";
                discourse.innerHTML += "topicId: " + data.post.Discourse + "};";
                discourse.innerHTML += "(function() {";
                discourse.innerHTML += "var d = document.createElement('script'); d.type = 'text/javascript'; d.async = true;";
                discourse.innerHTML += "d.src = DiscourseEmbed.discourseUrl + 'javascripts/embed.js';";
                discourse.innerHTML += "(document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(d);";
                discourse.innerHTML += "})();";
                discourse.async = true;
                document.body.appendChild(discourse);
            })
            .catch(console.log);
    }

    render() {
        if (this.state.post) {
            return (
                <div className="page-content">
                    <div className="wrap yue">
                        <div className="post">
                            <header className="post-header yue">
                                <h1>{ this.state.post.Title }</h1>
                            </header>
                            <article className="post-content">
                                <div dangerouslySetInnerHTML={
                                    {__html: this.state.post.Content}
                                } className="yue">
                                </div>
                            </article>
                        </div>
                        <div id='discourse-comments'></div>
                    </div>
                    <div className="wrap yue">
                        <span style={{ display: 'inline', float: 'left' }}>
                            <small>本页面访问量{ this.state.postView }次</small>
                        </span>
                        <span style={{ display: 'inline', float: 'right' }}>
                            <small>总访问量{ this.state.totalView }次</small>
                        </span>
                    </div>
                </div>
            );
        }
        return (
            <div></div>
        )
    }
}

export default Post;