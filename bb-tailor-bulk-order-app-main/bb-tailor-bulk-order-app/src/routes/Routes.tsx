import React from 'react';
import { Routes, Route, Navigate, Outlet } from 'react-router-dom';
import { Authenticated } from '@refinedev/core';
import { Login, Home } from '@pages';

const AppRoutes: React.FC = () => {
    return (
        <Routes>
            <Route
                element={
                    <Authenticated fallback={<Outlet />}>
                        <Navigate to="/home" />
                    </Authenticated>
                }
            >
                <Route path="/login" element={<Login />} />
            </Route>
            <Route
                element={
                    <Authenticated redirectOnFail="/login">
                        <Outlet />
                    </Authenticated>
                }
            >
                <Route path="/home" element={<Home />} />
                <Route
                    path="/"
                    element={<Navigate to="/home" replace/>}
                />
                {/* Add more routes here as needed */}
            </Route>

            {/* Add more routes here as needed */}
        </Routes>
    );
};

export default AppRoutes;