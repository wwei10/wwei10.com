import React from 'react';

class Post extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            permalink: this.props.match.params.permalink,
            post: null,
        };
    }

    componentDidMount() {
        fetch('http://localhost:8080/api/v1/search/' + this.state.permalink)
            .then(res => res.json())
            .then((data) => {
                this.setState({
                    post: data.post,
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