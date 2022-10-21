
import classes from "./Account.scss";
import {Link, useNavigate} from "react-router-dom";
import Cookies from "universal-cookie";
import {useEffect, useState} from "react";
import {toast} from "react-toastify";

const Account = ({token, setToken}) => {
    let navigate = useNavigate();
    let cookies = new Cookies()
    const baseEvent = {
        id: -1,
        name: "Самая лучшее мероприятие",
        description: "",
        short: "info",
        img: "",
        status: "123",
        holder_id: -1,
    }
    const [myEvents, setEventValue] = useState([]);
    const [myOlympiads, setOlympiadValue] = useState([]);
    const [isAdmin, setAdminValue] = useState(false)

    let handleQuit = (e) => {
        e.preventDefault()
        cookies.remove("token")
        setToken()
        navigate("/")
    }

    const getEvents = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        // console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        // console.log(process.env.REACT_APP_API_URL+"/events/my")
        fetch(process.env.REACT_APP_API_URL+"/events/my", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Origin': window.location.origin.toString(),
                'Authorization': "Bearer " + token,
            }
        })
            .then(response =>
            {if (response.status == 401) {
                cookies.remove("token")
                navigate("/authorize")
            }else{
                response.json().then(json=>setEventValue(json.events))}})
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    const getOlympiads = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        // console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        // console.log(process.env.REACT_APP_API_URL+"/olympiads/my")
        fetch(process.env.REACT_APP_API_URL+"/olympiads/my", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Origin': window.location.origin.toString(),
                'Authorization': "Bearer " + token,
            }
        })
            .then(response => response.json())
            .then(json =>{
                // console.log(json.olympiads)
                setOlympiadValue(json.olympiads)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }


    const checkAdmin = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        // console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        // console.log(process.env.REACT_APP_API_URL+"/admin/check")
        fetch(process.env.REACT_APP_API_URL+"/admin/check", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Origin': window.location.origin.toString(),
                'Authorization': "Bearer " + token,
            }
        })
            .then(response => {
                if (response.status === 200) {
                    setAdminValue(true)
                }
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    useEffect(() => {
        getEvents()
        getOlympiads()
        checkAdmin()
    }, [])

    return (isAdmin) ? (
        <div className="account-container">
            <div>
                <Link to="/admin" style={{backgroundColor: "red", color: "white",padding:"10px", borderRadius: "25px", display:"block", margin:"auto", maxWidth:"100px", textAlign:"center"}}>Админская панель</Link>
                <p className="Olympiad" style={{color:"black"}}>Мои олимпиады</p>
                <div className="myOlympiad">
                    {myOlympiads.map((item, index) => (
                        <div className="myOlympiadsOlympiad">
                            <div className="myOlympiadsOlympiadImg">
                                <img src={process.env.REACT_APP_STATIC_FILES + "/" + item.img}></img>
                            </div>
                            <Link className="OlympiadLink" to={"/olympiad/" + item.big_olympiad_id + "/" + item.id}>{item.name}</Link>
                        </div>
                    ))}
                </div>
                <p className="Event" style={{color:"black"}}>Мои мероприятия</p>
            <div className="myEvent">
                {myEvents.map((item, index) => (
                    <div className="myEventsEvent">
                        <div className="myEventsEventImg">
                            <img src={process.env.REACT_APP_STATIC_FILES + "/" + item.img}></img>
                        </div>
                        <Link className="EventLink" to={"/event/" + item.id}>{item.name}</Link>
                    </div>
                ))}
            </div>
                <div>
                    <a href="#" style={{backgroundColor: "red", color: "white",padding:"10px", borderRadius: "25px", display:"block", margin:"auto", maxWidth:"60px", textAlign:"center"}} onClick={handleQuit}>Выйти</a>
                </div>
            </div>
        </div>


    ) : (
        <div className="account-container">
            <div>
                <p className="Olympiad" style={{color:"black"}}>Мои олимпиады</p>
                <div className="myOlympiad">
                    {myOlympiads.map((item, index) => (
                        <div className="myOlympiadsOlympiad">
                            <div className="myOlympiadsOlympiadImg">
                                <img src={process.env.REACT_APP_STATIC_FILES + "/" + item.img}></img>
                            </div>
                            <Link className="OlympiadLink" to={"/olympiad/" + item.big_olympiad_id + "/" + item.id}>{item.name}</Link>
                        </div>
                    ))}
                </div>
                <p className="Event" style={{color:"black"}}>Мои мероприятия</p>
                <div className="myEvent">
                    {myEvents.map((item, index) => (
                        <div className="myEventsEvent">
                            <div className="myEventsEventImg">
                                <img src={process.env.REACT_APP_STATIC_FILES + "/" + item.img}></img>
                            </div>
                            <Link className="EventLink" to={"/event/" + item.id}>{item.name}</Link>
                        </div>
                    ))}
                </div>
                <div>
                    <a href="#" style={{backgroundColor: "red", color: "white",padding:"10px", borderRadius: "25px", display:"block", margin:"auto", maxWidth:"60px", textAlign:"center"}} onClick={handleQuit}>Выйти</a>
                </div>
            </div>
        </div>
    )
}
export default Account;

