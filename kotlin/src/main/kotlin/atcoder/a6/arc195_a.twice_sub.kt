package atcoder.a6

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val (N, M) = input_ints()
    val A = input_ints()
    val B = input_ints()

    val left = mutableListOf<Int>()
    var i = 0
    for(b in B){
        i = (i until N).find{A[it] == b} ?: N
        if(i >= N){ break }
        left.add(i)
        i += 1
    }
    if(left.size != M){
        println("No")
        return
    }

    val right = mutableListOf<Int>()
    i = N-1
    for(b in B.reversed()){
        i = (i downTo 0).find{A[it] == b} ?: -1
        if(i < 0){
            break
        }
        right.add(i)
        i -= 1
    }
    println(if(right.reversed().zip(left).any{(a,b) -> a!=b}) "Yes" else "No")
}
