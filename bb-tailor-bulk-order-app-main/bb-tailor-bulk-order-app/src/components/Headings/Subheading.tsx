import React from 'react'

interface SubheadingProps {
    className?: string
    text?: string
}

const Subheading: React.FC<SubheadingProps> = ({ className, text }) => {
    return <h2 className={`font-nunito text-subheading text-secondaryText ${className ? className : ''}`}> {text} </h2>
}

export default Subheading
