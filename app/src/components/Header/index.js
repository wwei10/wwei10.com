import React from 'react';

class Header extends React.Component {
    render() {
        return (
            <header className="site-header">
                <div className="wrap">
                    <a className="site-title" href="/">Wei's Blog</a>
                    <nav className="site-nav">
                        <div className="trigger">
                            <a className="page-link yue" href="/about">关于</a>
                        </div>
                    </nav>
                </div>
            </header>
        );
    }
}

export default Header;