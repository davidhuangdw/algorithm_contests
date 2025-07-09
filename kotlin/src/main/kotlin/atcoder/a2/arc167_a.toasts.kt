package atcoder.a2

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    val (N, M) = input_ints()
    val A = input_ints().sorted()
    var sum = A.fold(0L){acc, v -> acc + 1L*v*v}
    val k = N - M
    for((x,y) in A.subList(0,k).reversed().zip(A.subList(k,k+k))) {
        sum += 2L*x*y
    }
    println(sum)
}