import React from 'react'

interface HeadingProps {
    className?: string
    text?: string
}

const Heading: React.FC<HeadingProps> = ({ className, text }) => {
    return <h1 className={`font-nunito text-heading text-primaryText ${className ? className : ''}`}> {text} </h1>
}

export default Heading
