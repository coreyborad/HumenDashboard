import React from "react";
import { TextField, MenuItem } from '@mui/material';
const payerList = [
    {
        value: "魚飼料",
        label: "魚飼料",
    },
    {
        value: "小魚",
        label: "小魚",
    },
    {
        value: "白菜",
        label: "白菜",
    }
]

export const Payer = (props: { children?: any, value: string, onChange: any }) => {
    const [payerState, setPayerState] = React.useState(props.value);
    return (
        <TextField
            fullWidth
            sx={{ mt: 3, mb: 2 }}
            margin="normal"
            id="outlined-select-currency"
            select
            label="付款人"
            value={payerState}
            onChange={(e) => {
                setPayerState(e.target.value)
                props.onChange(e)
            }}
        >
            {payerList.map((option) => (
                <MenuItem key={option.value} value={option.value}>
                    {option.label}
                </MenuItem>
            ))}
        </TextField>
    );
}