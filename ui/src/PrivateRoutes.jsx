import { Navigate, Outlet } from "react-router-dom"

const PrivateRoute = ({user}) => {
    return(
        Object.keys(user).length !== 0 ? <Outlet /> : <Navigate to ="/login" />
    )
}

export default PrivateRoute