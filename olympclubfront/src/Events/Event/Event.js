import {Link, useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import classes from "./Event.scss";
import {toast, ToastContainer} from "react-toastify";
import md5 from "blueimp-md5";

const Event = ({token}) => {
    const params = useParams();
    const baseEvent = {
        id: -1,
        name: "Самая лучшее мероприятие",
        description: "",
        short: "info",
        img: "",
        status: "123",
        holder_id: -1,
    }
    const [event, setValue] = useState(baseEvent);
    const [news, setNews] = useState([])
    const eventID = params.event

    const getNews = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        fetch(process.env.REACT_APP_API_URL+"/event/" + eventID + "/news", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json =>{
                setNews(json.news)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    const getHolder = (event) => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        fetch(process.env.REACT_APP_API_URL+"/holder/" + event.holder_id, {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json =>{
                event.holder = json.holder
                setValue(event)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }


    const getEvent = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        fetch(process.env.REACT_APP_API_URL+"/event/" + eventID, {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json =>{
                getHolder(json.event)
                })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    useEffect(() => {
        getNews()
        getEvent()
    }, [])



    const subscribeEvent = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        headers.append('Authorization', "Bearer " + token)
        fetch(process.env.REACT_APP_API_URL+"/event/" + eventID + "/add", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => {
                if (response.status !== 200) {
                    response.json().then((errorJson) => {
                        // console.log(errorJson.error)
                        toast(errorJson.error)
                    });
                }else{
                    toast("Connection is created")
                }}).catch(error => console.log('Authorization failed: ' + error.message));
    }

    return (event.holder) ? (
        <div className="container">
            <ToastContainer/>
            <table className={`${classes.olympiad} olympiad`}>
                <tr>
                    <td colSpan={2}>
                        <Link className="OlympiadLink" to={"/event/" + event.id}>{event.name}</Link>
                    </td>
                    <td>
                        <p dangerouslySetInnerHTML={{__html: event.status}}></p>
                    </td>
                </tr>
                <tr>
                    <td rowSpan={2}>
                        <div className={`${classes.olympiadImg} olympiadImg`}>
                            <img src={process.env.REACT_APP_STATIC_FILES + "/" + event.img}></img>
                        </div>
                    </td>
                    <td style={{paddingLeft: "20px"}}>
                        <p>Организатор</p>
                        <div style={{display: "flex"}}>
                            <img style={{ width: "50px", height: "50px"}} src={process.env.REACT_APP_STATIC_FILES + "/" + event.holder.logo}></img>
                            <p style={{marginLeft: "20px"}}>{event.holder.name}</p>
                        </div>
                    </td>
                    <td>
                        <p> Сайт: <a href={event.website}>link</a></p>
                    </td>
                </tr>
                <tr>
                    <td>
                    </td>
                    <td>
                        <a className={`${classes.ButtonLink} ButtonLink`} onClick={subscribeEvent}> Отслеживать Олимпиаду </a>
                    </td>
                </tr>
            </table>
            <p style={{fontWeight: "bold", fontSize: "32px"}}>Новости</p>
            <div className="EventsNews">
                {news.map((item, index) => (
                    <div className="EventsNewsNews">
                        <p className="EventNews">{item.Title}</p>
                        <p style={{textAlign: "center"}}>{item.Description}</p>
                    </div>
                ))}
            </div>
        </div>
    ) : (
        <div>
            <p>{event.name}</p>
            <p>{event.img}</p>
        </div>
    )
}

export default Event;


