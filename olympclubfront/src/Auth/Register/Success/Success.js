import {Route, useLocation, useParams, Routes, Link} from "react-router-dom";
import classes from "./Success.scss";
import {useEffect, useState} from "react";

const RegisterSuccess = () => {
    return (
            <div className="change-password-container">
                <div className="change-password-window">
                    <p className="change-password-title">Подтверждение почты</p>
                    <p>Письмо успешно отправлено на вашу почту :)</p>
                    <p>Пройдите по ссылке, чтобы подтвердить почту</p>
                    <div className="img-container">
                        <Link to="/">
                            <img src={process.env.REACT_APP_STATIC_FILES + "/icons/mail_icon.png"}/>
                        </Link>
                    </div>
                </div>
            </div>
    )
}

export default RegisterSuccess