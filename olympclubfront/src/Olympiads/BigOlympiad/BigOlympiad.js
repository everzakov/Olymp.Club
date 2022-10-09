import {Route, useLocation, useParams, Routes, Link} from "react-router-dom";
import Olympiad from "../Olympiad/Olympiad";
import {useEffect, useState} from "react";
import classes from "./BigOlympiad.scss";

const BigOlympiad = ({token}) => {
    console.log("token", token)
    let match = useLocation();
    const params = useParams();
    console.log(params.bigOlympiad)
    console.log(match.url)
    const [olympiad, setValue] = useState(0);

    const baseBigOlympiad = {
        id: -1,
        name: "Самая лучшая олимпиада",
        short: "",
        logo: "",
        description: "",
        olympiads: [],
    }
    const [bigOlympiad, setBigOlympiadValue] = useState(baseBigOlympiad)

    const getOlympiads = (bigOlympiad) => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        console.log(process.env.REACT_APP_API_URL+"/olympiad/" + params.bigOlympiad + "/olympiads")
        fetch(process.env.REACT_APP_API_URL+"/olympiad/" +  params.bigOlympiad + "/olympiads", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json =>{
                console.log(json.olympiads)
                bigOlympiad.olympiads = json.olympiads
                setBigOlympiadValue(bigOlympiad)

            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    const getOlympiad = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        console.log(process.env.REACT_APP_API_URL+"/olympiad/" + params.bigOlympiad)
        fetch(process.env.REACT_APP_API_URL+"/olympiad/" +  params.bigOlympiad, {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json =>{
                getOlympiads(json.big_olympiad)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    const reRender = (e) => {
      setValue(olympiad + 1)
    }

    useEffect(() => {
        getOlympiad()
    }, [])

    return (
        <div>
            <div className={`${classes.olympiadContainer} olympiadContainer`}>
                    <table className={`${classes.olympiad} bigOlympiad`}>
                        <tr>
                            <td>
                                <p className="BigOlympiadLink">{bigOlympiad.name}</p>
                            </td>
                            <td>
                            </td>
                            <td>
                                <p dangerouslySetInnerHTML={{__html: bigOlympiad.status}}></p>
                            </td>
                        </tr>
                        <tr>
                            <td rowSpan={2}>
                                <div className={`${classes.olympiadImg} olympiadImg`}>
                                    <img src={process.env.REACT_APP_STATIC_FILES + "/" + bigOlympiad.logo}></img>
                                </div>
                            </td>
                            <td>
                                <p>{bigOlympiad.description}</p>
                            </td>
                            <td>
                            </td>
                        </tr>
                        <tr>
                            <td>
                            </td>
                            <td>
                            </td>
                        </tr>
                    </table>
                <div className="bigOlympiadolympiads">
                    {bigOlympiad.olympiads.map((item, index) => (
                        <div className="bigOlympiadolympiadsOlympiad">
                            <div className="bigOlympiadolympiadsImg">
                            <img src={process.env.REACT_APP_STATIC_FILES + "/" + item.img}></img>
                            </div>
                            <Link onClink={reRender} className="OlympiadLink" to={"/olympiad/" + bigOlympiad.id + "/" + item.id}>{item.name}</Link>
                        </div>
                    ))}
                </div>
            </div>
            <div>
            <Routes>
                <Route path={`:olympiad`} element={<Olympiad token={token} />} />
            </Routes>
            </div>
        </div>
    )
}

export default BigOlympiad