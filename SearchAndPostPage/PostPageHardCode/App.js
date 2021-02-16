import TopBar from "./TopBar";
import {Layout} from "antd";
import Aside from "./Aside";
import Main from "./Main";
import Crumb from "./Crumb";

function App() {
  return (
    <div className="App">
        <Layout>
            <TopBar />
            <Layout>
                <Aside/>
                <Layout style={{ padding: '0 24px 24px' }}>
                    <Crumb/>
                    <Main/>
                </Layout>
            </Layout>
        </Layout>

    </div>
  );
}

export default App;
