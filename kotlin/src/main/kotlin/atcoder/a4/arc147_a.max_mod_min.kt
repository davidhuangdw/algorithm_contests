package atcoder.a4

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}

    val (N, ) = input_ints()
    val A = ArrayDeque(input_ints().sorted())
    var cnt = 0
    while(A.size > 1){
        cnt++
        val r = A.removeLast() % A.first()
        if(r > 0){ A.addFirst(r)}
    }
    println(cnt)
}
