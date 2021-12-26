import React from "react";
import { TextField } from '@mui/material';

export const Note = (props: { children?: any, value: string, onChange: any }) => {
    const [noteState, setNoteState] = React.useState(props.value);
    return (
        <TextField
            fullWidth
            sx={{ mt: 3, mb: 2 }}
            margin="normal"
            id="outlined-select-currency"
            label="項目"
            value={noteState}
            onChange={(e) => {
                setNoteState(e.target.value)
                props.onChange(e)
            }}
        >
        </TextField>
    );
}