package main

func main() {
	println(parseBoolExpr("!(&(!(&(f)),&(t),|(f,f,t)))"))
}
