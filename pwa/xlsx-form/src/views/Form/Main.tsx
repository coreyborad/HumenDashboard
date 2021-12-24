import React from "react";

interface UserInfo {
    name: string;
    lastname: string;
}

export const Form: React.FC = () => {
    const [userInfo, setUserInfo] = React.useState<UserInfo>({
        name: "John",
        lastname: "Doe",
    });
    return (
        <div>
            <form>
                <h4>
                    {userInfo.name} {userInfo.lastname}
                </h4>
                <input
                    value={userInfo.name}
                    onChange={(e) =>
                        setUserInfo({
                            ...userInfo,
                            name: e.target.value,
                        })
                    }
                />
            </form>
        </div>
    )
}
