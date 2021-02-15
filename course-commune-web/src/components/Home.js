import TopBar from "./TopBar";
import {Layout} from "antd";
import Aside from "./Aside";
import Crumb from "./Crumb";
import MainContent from "./MainContent"
import Pinterest from "./tutor1"
import CardList from "./tutor2"

function Home() {
    return (
        <div className="App">
            <Layout>
                <Layout>
                    <Aside/>
                    {/*在这个layout 里切换组件*/}
                    <Layout style={{ padding: '0 24px 24px' }}>
                        <Crumb/>
                        <MainContent/>
                        {/*<Pinterest/>*/}
                        {/*<CardList/>*/}
                    </Layout>
                </Layout>
            </Layout>
        </div>
    );
}

export default Home;