import React from 'react';

class About extends React.Component {
    render() {
        return (
            <div className="page-content">
                <div className="post wrap yue">
                    <header className="post-header">
                        <h1>关于</h1>
                    </header>
                    <article className="post-content">
                        <p>我是魏炜，Instagram ads的一名工程师。之前有幸在Stanford，上海交大学习。</p>

                        <p>这里的内容会覆盖我编程学习的经历，职业成长的反思，生活琐事等等。争取自己每天能进步一点点！</p>

                        <p>个人百科的地址：<a href="https://docs.wwei10.com">docs.wwei10.com</a></p>

                        <hr />
                            <p>Hi all,</p>
                            <p>I am Wei, an engineer working on Instagram ads. I got a degree in MSCS at Stanford and got my bachelor in Shanghai Jiao Tong University.</p>
                            <p>I believe in 1% rule and hope to improve myself everyday. I will cover my improvements here and keep track of my progress. Stay tuned.</p>
                    </article>
                </div>
            </div>
        );
    }
}

export default About;