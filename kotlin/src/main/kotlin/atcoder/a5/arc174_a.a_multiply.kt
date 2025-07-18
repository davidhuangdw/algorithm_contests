package atcoder.a5

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
    val (N, C) = inputInts()
    val A = inputInts()
    var (mn, mx, sum, part_mn, part_mx) = LongArray(5)

    for(v in A){
        sum += v
        part_mn = minOf(part_mn, sum - mx)
        part_mx = maxOf(part_mx, sum - mn)
        mn = minOf(mn, sum)
        mx = maxOf(mx, sum)
    }
    println(sum+maxOf(0, part_mn*(C-1), part_mx*(C-1)))
}
