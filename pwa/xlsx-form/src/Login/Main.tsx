import React from "react";
import { Box, TextField, Container, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { login } from "../apis/User";
import { User } from "../interfaces/Interface"

export const Login: React.FC = () => {
    const navigate = useNavigate();  
    const [user, setUser] = React.useState<User>({
        email: "",
        password: "",
      });
    const handleLogin = async () => {
        await login(user);
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
                    value={user.email}
                    onChange={(e) =>
                        setUser({
                          ...user,
                          email: e.target.value,
                        })
                      }
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
                    value={user.password}
                    onChange={(e) =>
                        setUser({
                          ...user,
                          password: e.target.value,
                        })
                      }
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
