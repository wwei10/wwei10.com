'use strict';

const e = React.createElement;

class Posts extends React.Component {
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
            <ul class="posts">
                {this.state.posts.map(post => (
                    <li key={post.Permalink}>
                        <span class="post-date yue">{post.Date}</span>
                        <a class="yue" href={post.Permalink}>{post.Title}</a>
                    </li>
                ))}
            </ul>
        );
    }
}

const domContainer = document.querySelector('#posts');
ReactDOM.render(
    <Posts />,
    domContainer,
);