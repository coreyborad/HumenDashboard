import React from "react";
import { Box, TextField, Container, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';  

export const Login: React.FC = () => {
    const navigate = useNavigate();  
    const handleLogin = () => {
        navigate('/form');
    }

    return (
        <Container component="main" maxWidth="xs">
            <Box
                component="form"
                sx={{
                    marginTop: 8,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
                noValidate
                autoComplete="off"
            >
                <TextField
                    margin="normal"
                    required
                    fullWidth
                    id="email"
                    label="Email Address"
                    name="email"
                    autoComplete="email"
                    autoFocus
                />
                <TextField
                    margin="normal"
                    required
                    fullWidth
                    name="password"
                    label="Password"
                    type="password"
                    id="password"
                    autoComplete="current-password"
                />
                <Button
                    variant="contained"
                    fullWidth
                    sx={{ mt: 3, mb: 2 }}
                    onClick={handleLogin}
                >
                    Login
                </Button>
            </Box>
        </Container>
    )
}
