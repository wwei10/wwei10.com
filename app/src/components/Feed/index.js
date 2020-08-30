import React from 'react';

class Feed extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            posts: [],
        };
    }


    componentDidMount() {
        fetch('http://localhost:8080/api/v1/timeline')
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
                </div>
            </div>
        );
    }
}

export default Feed;