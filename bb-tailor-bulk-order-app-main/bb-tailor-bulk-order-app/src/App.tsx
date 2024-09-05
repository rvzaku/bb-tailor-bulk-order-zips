import AppRoutes from './routes/Routes'
import { Refine } from '@refinedev/core';
import { authProvider } from '@providers';
import { BrowserRouter } from "react-router-dom";
import routerProvider from "@refinedev/react-router-v6";

function App() {
    return (
        <BrowserRouter>
            <Refine 
                authProvider={authProvider} 
                routerProvider={routerProvider}
            >
                <AppRoutes />
            </Refine>
        </BrowserRouter>
    );
}

export default App
