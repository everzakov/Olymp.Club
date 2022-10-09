import {Link, Routes, Route, useLocation, useNavigate} from "react-router-dom";
import {useEffect, useState} from "react";
import {toast} from "react-toastify";
import classes from "./Admin.scss";
import Cookies from "universal-cookie";
import AddHolder from "./Holder/Holder";
import AddBigOlympiad from "./BigOlympiad/BigOlympiad";
import AddOlympiad from "./Olympiad/Olympiad";
import AddEvent from "./Event/Event";
import AddNews from "./News/News";
const AdminPanel = ({token}) => {
    const navigate = useNavigate();
    const cookies = new Cookies();

    const VerifyUser = () => {
        let headers = new Headers();
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        headers.append("Authorization", "Bearer " + token)
        fetch(process.env.REACT_APP_API_URL+"/admin/check", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => {
                console.log(response.status)
                if (response.status === 401) {
                    cookies.remove("token")
                    navigate("/")
                }
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    useEffect(() => {
        VerifyUser()
    }, [])

    return (

        <div>
        <div className="window-container">
            <div className="window">
                <Link className="window-button" to="/admin/holder">Add a holder</Link>
                <Link className="window-button" to="/admin/bigOlympiad">Add a bigOlympiad</Link>
                <Link className="window-button" to="/admin/olympiad">Add an olympiad</Link>
                <Link className="window-button" to="/admin/event">Add an event</Link>
                <Link className="window-button" to="/admin/news">Add news</Link>
            </div>
        </div>
                <Routes>
                    <Route path={"/holder"} element={<AddHolder token={token}/> }/>
                    <Route path={"/bigOlympiad"} element = {<AddBigOlympiad token={token}/> }/>
                    <Route path={"/olympiad"} element={<AddOlympiad token={token}/> }/>
                    <Route path={"/event"} element = {<AddEvent token={token}/> } />
                    <Route path={"/news"} element={<AddNews token={token}/> }/>
                </Routes>
        </div>
    )
}

export default AdminPanel