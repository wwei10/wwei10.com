import React from 'react';

class Post extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            post: null,
        };
    }

    componentDidMount() {
        fetch('http://localhost:8080/api/v1/timeline')
            .then(res => res.json())
            .then((data) => {
                this.setState({
                    post: data.posts[0],
                });
            })
            .catch(console.log);
    }

    render() {
        if (this.state.post) {
            return (
                <div class="page-content">
                    <div class="wrap yue">
                        <div class="post">
                            <header class="post-header">
                                <h1>{ this.state.post.Title }</h1>
                            </header>
                            <article class="post-content">
                                { this.state.post.Content }
                            </article>
                        </div>
                        <div id='discourse-comments'></div>
                    </div>
                </div>
            );
        }
        return (
            <div class="wrap yue">Loading ... </div>
        )
    }
}

export default Post;