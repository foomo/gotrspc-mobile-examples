import Image from "next/image";
import { Upload } from "lucide-react";

import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import Header from "@/components/Header";
import { ModeToggle } from "@/components/ModeToggle";

export default function Home() {
	return (
		<>
			<main className="flex min-h-screen flex-col gap-4 p-4 pt-[140px]">
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
				<Card className="overflow-hidden">
					<CardHeader>
						<CardTitle>image.png</CardTitle>
						<CardDescription>
							Lipsum dolor sit amet, consectetur adipiscing elit
						</CardDescription>
					</CardHeader>
				</Card>
			</main>
		</>
	);
}
