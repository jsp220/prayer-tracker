interface buttonProps {
    label: string;
    className: string | undefined;
}

export default function Button(props: buttonProps) {
    return <button className={props.className}>{props.label}</button>;
}
