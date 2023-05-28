import { useAuthState } from '@/store/useAuth';
import React from 'react';
import { Navigate, Outlet } from "react-router-dom";


const ProtectedRoute = ({ allowedRoles }) => {
  const user = useAuthState(state => state.user)
  return allowedRoles.find((role) => user?.role?.includes(role)) ? (
    <Outlet />
  ) : (
    <Navigate to='/login' replace />
  )
}

export default ProtectedRoute;