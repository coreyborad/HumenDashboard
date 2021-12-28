import React from "react";
import { Accounting } from "../../interfaces/Interface";
import { Box, TextField, Container, Button, TextFieldProps } from '@mui/material';
import { MobileDatePicker } from '@mui/lab';
import AdapterDayjs from '@mui/lab/AdapterDayjs';
import LocalizationProvider from '@mui/lab/LocalizationProvider';
import { Payer } from "./components/Payer";
import { Category } from "./components/Category";
import { Note } from "./components/Note"
import { Cost } from "./components/Cost"
import { appendRecord } from "../../apis/Xlsx"

export const Form: React.FC = () => {
    const [accountingInfo, setAccountingInfo] = React.useState<Accounting>({
        date: new Date(),
        payer: "魚飼料",
        category: "食",
        subCategory: "外食",
        note: "",
        cost: 0
    });

    const handleSubmit = async () => {
        await appendRecord(accountingInfo)
        alert("success")
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
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <MobileDatePicker
                        label="日期"
                        inputFormat="YYYY MM/DD"
                        value={accountingInfo.date}
                        onChange={(value: any) =>
                            setAccountingInfo({
                                ...accountingInfo,
                                date: value,
                            })
                        }
                        renderInput={(params: JSX.IntrinsicAttributes & TextFieldProps) => <TextField {...params} />}
                    />
                </LocalizationProvider>
                <Payer
                    value={accountingInfo.payer}
                    onChange={(
                        e: React.ChangeEvent<HTMLInputElement>,
                    ): void => setAccountingInfo({
                        ...accountingInfo,
                        payer: e.target.value,
                    })
                    }
                >
                </Payer>
                <Category
                    category={accountingInfo.category}
                    subCategory={accountingInfo.subCategory}
                    onCategoryChange={(
                        value: string,
                    ): void => setAccountingInfo({
                        ...accountingInfo,
                        category: value,
                    })
                    }
                    onSubCategoryChange={(
                        value: string,
                    ): void => setAccountingInfo({
                        ...accountingInfo,
                        subCategory: value,
                    })
                    }
                >

                </Category>
                <Note
                    value={accountingInfo.note}
                    onChange={(
                        e: React.ChangeEvent<HTMLInputElement>,
                    ): void => setAccountingInfo({
                        ...accountingInfo,
                        note: e.target.value,
                    })}
                >
                </Note>
                <Cost
                    value={accountingInfo.cost}
                    onChange={(
                        value: number,
                    ): void => setAccountingInfo({
                        ...accountingInfo,
                        cost: value,
                    })}
                >
                </Cost>
                
                <Button
                    variant="contained"
                    fullWidth
                    sx={{ mt: 3, mb: 2 }}
                    onClick={handleSubmit}
                >
                    Submit
                </Button>
            </Box>
        </Container>
    )
}
