import React, { useState } from 'react'
import { Logo, LinkButton, PrimaryButton, Input, Heading, Subheading } from '@components'
import { useScreenSize } from '@hooks'
import { useLogin } from '@refinedev/core'

const Login: React.FC = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const logoSize = useScreenSize('50px', '50px')

    const { mutate: login } = useLogin()

    const handleLogin = async () => {
        login({ email, password })
    }

    return (
        <>
            <Logo
                className="fixed top-4 left-4 sm:top-8 sm:left-8 md:top-8 md:left-8 lg:top-16 lg:left-16 xl:top-32 xl:left-32 2xl:top-32 2xl:left-32"
                width={logoSize.width}
                height={logoSize.width}
            />
            <div className="flex flex-col justify-center items-center h-screen">
                <div className="text-center space-y-4">
                    <Heading text="welcome back" />
                    <Subheading text="Brooks Bingham Tailor & Bulk Orders" />
                    <Input
                        type="text"
                        placeholder="email"
                        name="email"
                        className="mt-4 h-10"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <Input
                        type="password"
                        placeholder="password"
                        name="password"
                        className="h-10"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <div className="flex justify-end w-full">
                        <LinkButton className="h-4 mt-0"> forgot password? </LinkButton>
                    </div>
                    <div className="flex w-full">
                        <PrimaryButton className="w-full h-10 mt-8" onClick={handleLogin}>
                            {' '}
                            log in{' '}
                        </PrimaryButton>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Login
