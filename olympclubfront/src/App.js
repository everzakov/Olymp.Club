import {React, useState} from "react";
import './App.css';
import Header from "./basic/header";
import {BrowserRouter, Routes, Route} from "react-router-dom";
import Olympiads from "./Olympiads/Olympiads/Olympiads";
import Olympiad from "./Olympiads/Olympiad/Olympiad"
import RegisterSuccess from "./Auth/Register/Success/Success";
import BigOlympiad from  "./Olympiads/BigOlympiad/BigOlympiad"
import Auth from  "./Auth/Auth/Auth"
import Register from "./Auth/Register/Register";
import RegisterVerify from "./Auth/Register/Verify/Verify"
import Account from "./Account/Account";
import Cookies from "universal-cookie";
import RegisterDecline from "./Auth/Register/Decline/Decline";
import PasswordRequest from "./Auth/Auth/Request/Request";
import PasswordRequestSuccess from "./Auth/Auth/Request/Success/Success"
import PasswordChange from "./Auth/Auth/Change/Change";
import PasswordChangeSuccess from "./Auth/Auth/Change/Success/Success";
import Events from "./Events/Events/Events";
import Event from "./Events/Event/Event";
import AdminPanel from "./Admins/Admin";
const App = () => {
    const cookies = new Cookies();
    const [token, setToken] = useState(cookies.get('token'))

    const SetTokenFunction = (token) => {
        setToken(token)
    }

    const DeleteTokenFunction = () => {
        setToken(undefined)
    }

    return (
        <BrowserRouter>
            <Header token={token}/>
            <Routes>
                <Route path="/me" element={<Account setToken={DeleteTokenFunction} token={token}/>}/>
                <Route path="/events" element={<Events token={token} />}/>
                <Route path="/event/:event" element={<Event token={token} />}/>
                <Route path="/admin/*" element={<AdminPanel token={token}/>} />
                <Route path="/olympiads" element={<Olympiads token={token} />}/>
                <Route path="/olympiad/:bigOlympiad/*" element={<BigOlympiad token={token}/>} />
                <Route path="/authorize/" element={<Auth setToken={SetTokenFunction}/>}/>
                <Route path="/register/" element={<Register />}/>
                <Route path="/register/success" element={<RegisterSuccess/>}/>
                <Route path="/register/verify" element={<RegisterVerify/>}/>
                <Route path="/register/decline" element={<RegisterDecline/>}/>
                <Route path="/password/request" element={<PasswordRequest/>}/>
                <Route path="/password/request/success" element={<PasswordRequestSuccess/>}/>
                <Route path="/password/change" element={<PasswordChange/>}/>
                <Route path="/password/change/success" element={<PasswordChangeSuccess/>}/>
            </Routes>
        </BrowserRouter>
    )
}

export default App;