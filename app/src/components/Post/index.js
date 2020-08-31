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
                    post: data.posts[6],
                });
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
                    </div>
                </div>
            );
        }
        return (
            <div className="wrap yue">Loading ... </div>
        )
    }
}

export default Post;