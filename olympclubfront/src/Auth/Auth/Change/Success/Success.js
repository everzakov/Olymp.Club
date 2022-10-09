import {Route, useLocation, useParams, Routes, Link} from "react-router-dom";
import classes from "./Success.scss";
import {useEffect, useState} from "react";

const PasswordChangeSuccess = () => {
    return (
            <div className="change-password-container">
                <div className="change-password-window">
                    <p className="change-password-title">Смена пароля</p>
                    <p>Вы успешно сменили пароль :)</p>
                    <Link className="window-button" to="/authorize">Авторизоваться</Link>
                </div>
            </div>
    )
}

export default PasswordChangeSuccess