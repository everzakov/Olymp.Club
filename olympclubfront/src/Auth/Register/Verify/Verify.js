import {Link, useLocation} from "react-router-dom";
import {useEffect, useState} from "react";
import {toast} from "react-toastify";
import classes from "./Verify.scss";

const RegisterVerify = () => {
    const search = useLocation().search;
    const token1 = new URLSearchParams(search).get('token1');
    const token2 = new URLSearchParams(search).get('token2');
    const [reason, setReason] = useState("")
    const [ok, setOk] = useState(false)


    const VerifyUser = () => {
        let headers = new Headers();
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        fetch(process.env.REACT_APP_API_URL+"/register/verify?" + "token1=" + token1 + "&token2=" + token2, {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => {
                if (response.status != 200) {
                    response.json().then((errorJson) => {
                        setReason(errorJson.error)
                    });
                }else{
                    setOk(true)
                }
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    useEffect(() => {
        VerifyUser()
    }, [])

    return (ok) ? (
        <div className="window-container">
            <div className="window">
                <h1 className="window-title">Регистрация подтверждена</h1>
                <p className="window-text">Вы успешно подтвердили регистрацию,<br/><b>авторизуйтесь</b>, чтобы продолжить
                    регистрацию</p>
                <Link className="window-button" to="/authorize">Авторизоваться</Link>
            </div>
        </div>
    ) : (
        <div>
            <div className="window-container">
                <div className="window">
                    <h1 className="window-title">Регистрация <b>не</b> подтверждена</h1>
                    <p className="window-text">По причине: {reason}</p>
                </div>
            </div>
        </div>
    )
}

export default RegisterVerify