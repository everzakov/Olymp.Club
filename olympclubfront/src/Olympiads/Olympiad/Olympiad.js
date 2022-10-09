import {Link, useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import classes from "./Olympiad.scss";
import {toast, ToastContainer} from "react-toastify";

const Olympiad = ({token}) => {
    const params = useParams();
    const baseOlympiad = {
        id: -1,
        name: "Самая лучшая олимпиада",
        subject: "info",
        level: "1level",
        img: "",
        short: "info",
        big_olympiad_id: -1,
        status: "123",
    }
    const [olympiad, setValue] = useState(baseOlympiad);
    const [news, setNews] = useState([])
    const bigOlympiadID = params.bigOlympiad
    const olympiadID = params.olympiad
    console.log(bigOlympiadID, olympiadID)

    const getNews = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        console.log(process.env.REACT_APP_API_URL+"/olympiad/" + bigOlympiadID + "/" + olympiadID + "/news")
        fetch(process.env.REACT_APP_API_URL+"/olympiad/" +  bigOlympiadID + "/" + olympiadID + "/news", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json =>{
                console.log(json.news)
                setNews(json.news)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    const getHolder = (olympiad) => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        console.log(process.env.REACT_APP_API_URL+"/holder/" + olympiad.holder_id)
        fetch(process.env.REACT_APP_API_URL+"/holder/" + olympiad.holder_id, {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json =>{
                console.log(json)
                olympiad.holder = json.holder
                setValue(olympiad)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }
    const getOlympiad = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        console.log(process.env.REACT_APP_API_URL+"/olympiad/" + bigOlympiadID + "/" + olympiadID)
        fetch(process.env.REACT_APP_API_URL+"/olympiad/" +  bigOlympiadID + "/" + olympiadID, {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json =>{
                console.log(json)
                getHolder(json.olympiad)
                })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    const subscribeOlympiad = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        console.log(token)
        headers.append('Authorization', "Bearer " + token)
        console.log(process.env.REACT_APP_API_URL+"/olympiad/" + bigOlympiadID + "/" + olympiadID + "/add")
        fetch(process.env.REACT_APP_API_URL+"/olympiad/" +  bigOlympiadID + "/" + olympiadID + "/add", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => {
            if (response.status !== 200) {
            response.json().then((errorJson) => {
                console.log(errorJson.error)
                toast(errorJson.error)
            });
        }else{
            toast("Connection is created")
        }}).catch(error => console.log('Authorization failed: ' + error.message));
    }

    useEffect(() => {
        getNews()
        getOlympiad()
    }, [olympiadID])


    return (olympiad.holder) ? (
        <div className="container">
            <ToastContainer/>
            <table className={`${classes.olympiad} olympiad`}>
                <tr>
                    <td>
                    </td>
                    <td>
                        <Link className="OlympiadLink" to={"/olympiad/" + olympiad.big_olympiad_id + "/" + olympiad.id}>{olympiad.name}</Link>
                    </td>
                    <td>
                        <p dangerouslySetInnerHTML={{__html: olympiad.status}}></p>
                    </td>
                </tr>
                <tr>
                    <td rowSpan={2}>
                        <div className={`${classes.olympiadImg} olympiadImg`}>
                            <img src={process.env.REACT_APP_STATIC_FILES + "/" + olympiad.img}></img>
                        </div>
                    </td>
                    <td style={{paddingLeft: "20px"}}>
                        <p>Организатор</p>
                        <div style={{display: "flex"}}>
                            <img style={{ width: "50px", height: "50px"}} src={process.env.REACT_APP_STATIC_FILES + "/" + olympiad.holder.logo}></img>
                            <p style={{marginLeft: "20px"}}>{olympiad.holder.name}</p>
                        </div>
                    </td>
                    <td>
                        <p> Сайт: <a href={olympiad.website}>link</a></p>
                    </td>
                </tr>
                <tr>
                    <td>
                    </td>
                    <td>
                        <a className={`${classes.ButtonLink} ButtonLink`} onClick={subscribeOlympiad}> Отслеживать Олимпиаду </a>
                    </td>
                </tr>
            </table>
            <p style={{fontWeight: "bold", fontSize: "32px"}}>Новости</p>
            <div className="OlympiadsNews">
                {news.map((item, index) => (
                    <div className="OlympiadsNewsNews">
                        <p className="OlympiadNews">{item.Title}</p>
                        <p style={{textAlign: "center"}}>{item.Description}</p>
                    </div>
                ))}
            </div>
        </div>
    ) : (
        <div>
            <p>{olympiad.subject}</p>
            <p>{olympiad.img}</p>
        </div>
    )
}

export default Olympiad;