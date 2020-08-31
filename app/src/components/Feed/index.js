import React from 'react';

class Feed extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            posts: [],
        };
    }


    componentDidMount() {
        var api = '/api/v1/timeline';
        if (window.location.hostname === 'localhost') {
            api = 'http://localhost:8080/api/v1/timeline';
        }
        fetch(api)
            .then(res => res.json())
            .then((data) => {
                this.setState({
                    posts: data.posts,
                });
            })
            .catch(console.log);
    }

    render() {
        return (
            <div className="page-content">
                <div className="wrap">
                    <ul className="posts">
                        {this.state.posts.map(post => (
                            <li key={post.Permalink}>
                                <span className="post-date yue">{post.Date}</span>
                                <a className="yue" href={post.Permalink}>{post.Title}</a>
                            </li>
                        ))}
                    </ul>
                    <p class="yue">通过
                        <a href="https://feedburner.google.com/fb/a/mailverify?uri=wwei10/blog&amp;loc=en_US">邮箱</a>
                        订阅
                    </p>
                </div>
            </div>
        );
    }
}

export default Feed;