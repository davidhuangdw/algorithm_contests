package atcoder.a4

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    val (_, A, B, C, D) = input_ints()
    val d = B - C
    println(if(-1 <=d && d<=1 && !(B==0 && C==0 && (A>0 && D>0))) "Yes" else "No")
}
