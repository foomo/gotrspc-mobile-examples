import { cn } from "@/lib/utils";
import { ModeToggle } from "./ModeToggle";

type HeaderProps = {
	className?: string;
};

const Header = (props: HeaderProps) => {
	return (
		<header
			className={cn(
				props.className,
				"header fixed z-20 left-0 right-0 top-0 w-full bg-background/90 border-b",
			)}
			style={{
				marginTop: "calc( env(safe-area-inset-top) * -1)",
			}}
		>
			<div />
			<nav
				className="p-4 backdrop-blur-xl"
				style={{
					paddingTop: "calc( env(safe-area-inset-top) * 2)",
				}}
			>
				<ModeToggle />
			</nav>
		</header>
	);
};

export default Header;
