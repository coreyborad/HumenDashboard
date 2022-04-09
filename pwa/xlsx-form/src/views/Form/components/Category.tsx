import React from "react";
import { TextField, MenuItem } from '@mui/material';
const categoryList = [
    {
        value: "基本支出",
        label: "基本支出",
        subCategory: [
            {
                value: "管理費",
                label: "管理費",
            },
            {
                value: "水費",
                label: "水費",
            },
            {
                value: "電費",
                label: "電費",
            },
            {
                value: "瓦斯費",
                label: "瓦斯費",
            },
            {
                value: "網路費",
                label: "網路費",
            },
            {
                value: "影音串流",
                label: "影音串流",
            }
        ]
    },
    {
        value: "食",
        label: "食",
        subCategory: [
            {
                value: "食材",
                label: "食材",
            },
            {
                value: "外食",
                label: "外食",
            }
        ]
    },
    {
        value: "生活",
        label: "生活",
        subCategory: [
            {
                value: "洗衣用品",
                label: "洗衣用品",
            },
            {
                value: "沐浴用品",
                label: "沐浴用品",
            },
            {
                value: "洗髮用品",
                label: "洗髮用品",
            },
            {
                value: "牙齒用品",
                label: "牙齒用品",
            },
            {
                value: "衛生紙、紙巾",
                label: "衛生紙、紙巾",
            },
            {
                value: "飲用水耗材",
                label: "飲用水耗材",
            },
            {
                value: "家事清潔",
                label: "家事清潔",
            },
            {
                value: "神秘物品",
                label: "神秘物品",
            },
            {
                value: "停車費、交通費",
                label: "停車費、交通費",
            },
            {
                value: "雜支",
                label: "雜支",
            }
        ]
    },
    {
        value: "國內旅遊",
        label: "國內旅遊",
        subCategory: [
            {
                value: "國內食",
                label: "國內食",
            },
            {
                value: "國內宿",
                label: "國內宿",
            },
            {
                value: "國內交通",
                label: "國內交通",
            },
            {
                value: "國內活動門票",
                label: "國內活動門票",
            }
        ]
    },
    {
        value: "國外旅遊",
        label: "國外旅遊",
        subCategory: [
            {
                value: "國外食",
                label: "國外食",
            },
            {
                value: "國外宿",
                label: "國外宿",
            },
            {
                value: "國外交通",
                label: "國外交通",
            },
            {
                value: "國外活動門票",
                label: "國外活動門票",
            }
        ]
    },
    {
        value: "娛樂",
        label: "娛樂",
        subCategory: [
            {
                value: "聚餐",
                label: "聚餐",
            },
            {
                value: "看電影",
                label: "看電影",
            },
            {
                value: "%%",
                label: "%%",
            }
        ]
    },
    {
        value: "樂透",
        label: "樂透",
        subCategory: [
            {
                value: "其他",
                label: "其他",
            }
        ]
    },
    {
        value: "稅",
        label: "稅",
        subCategory: [
            {
                value: "所得稅",
                label: "所得稅",
            },
            {
                value: "房屋稅",
                label: "房屋稅",
            },
            {
                value: "牌照稅",
                label: "牌照稅",
            },
            {
                value: "燃料稅",
                label: "燃料稅",
            },
            {
                value: "地價稅",
                label: "地價稅",
            }
        ]
    },
    {
        value: "修繕",
        label: "修繕",
        subCategory: [
            {
                value: "家具",
                label: "家具",
            },
            {
                value: "家電",
                label: "家電",
            },
            {
                value: "家飾收納",
                label: "家飾收納",
            },
            {
                value: "修繕",
                label: "修繕",
            }
        ]
    },
    {
        value: "醫療",
        label: "醫療",
        subCategory: [
            {
                value: "產檢",
                label: "產檢",
            },
            {
                value: "保健食品",
                label: "保健食品",
            },
            {
                value: "一般生病",
                label: "一般生病",
            }
        ]
    }
]

export const Category = (props: { children?: any, category: string, subCategory: string, onCategoryChange: any, onSubCategoryChange: any }) => {
    const [categoryState, setCategoryState] = React.useState(props.category);
    const [subCategoryState, setSubCategoryState] = React.useState(props.subCategory);
    const [selectedCategory, setSelectedCategory] = React.useState(categoryList.find( (item) => {
        return item.value === categoryState
    }));
    
    React.useEffect(() => {
        // Prevent the subCategory value is slow than selectCategory set
        const thisSelectedCategory = categoryList.find( (item) => {
            return item.value === categoryState
        })
        if(thisSelectedCategory?.subCategory){
            const subCategoryValue = thisSelectedCategory?.subCategory[0].value
            setSubCategoryState(subCategoryValue)
            props.onSubCategoryChange(subCategoryValue)
        }
        // To Set selectCategory generate menuItems
        setSelectedCategory(categoryList.find( (item) => {
            return item.value === categoryState
        }))

    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [categoryState])
    return (
        <React.Fragment>
            <TextField
                fullWidth
                sx={{ mt: 3, mb: 2 }}
                margin="normal"
                id="outlined-select-currency"
                select
                label="分類"
                value={categoryState}
                onChange={(e) => {
                    setCategoryState(e.target.value)
                    props.onCategoryChange(e.target.value)

                }}
            >
                {categoryList.map((option) => (
                    <MenuItem key={option.value} value={option.value}>
                        {option.label}
                    </MenuItem>
                ))}
            </TextField>
            <TextField
                fullWidth
                sx={{ mt: 3, mb: 2 }}
                margin="normal"
                id="outlined-select-currency"
                select
                label="子類"
                value={subCategoryState}
                onChange={(e) => {
                    setSubCategoryState(e.target.value)
                    props.onSubCategoryChange(e.target.value)
                }}
            >
                {selectedCategory?.subCategory.map((option) => (
                    <MenuItem key={option.value} value={option.value}>
                        {option.label}
                    </MenuItem>
                ))}
            </TextField>
        </React.Fragment>

    );
}