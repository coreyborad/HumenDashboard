import React from "react";
import { TextField, InputAdornment } from '@mui/material';

export const Cost = (props: { children?: any, value: number, onChange: any }) => {
    const [costState, setCostState] = React.useState(props.value);
    return (
        <TextField
            fullWidth
            sx={{ mt: 3, mb: 2 }}
            InputProps={{
                startAdornment: <InputAdornment position="start">$</InputAdornment>,
            }}
            margin="normal"
            label="花費"
            type="tel"
            value={costState}
            onChange={(e) => {
                const cost = Number(e.target.value)
                setCostState(cost)
                props.onChange(cost)
            }}
        >
        </TextField>
    );
}