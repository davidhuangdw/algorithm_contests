package atcoder.a5

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
    fun inputLongs() = readLine()!!.split(' ').map(String::toLong)

    val (T) = inputInts()
    repeat(T) {
        var (n2, n3, n4) = inputLongs()
        var cnt = 0L

        var k = minOf(n4+n2/2, n3/2)
        cnt += k
        n2 -= 2* maxOf(k-n4, 0)
        n4 -= minOf(k, n4)
//        println(listOf(cnt, n2,n3,n4))

        k = minOf(n4/2, n2)
        cnt += k
        n2 -= k
        n4 -= k*2
//        println(listOf(cnt, n2,n3,n4))

        k = minOf(n4, n2/3)
        cnt += k
        n2 -= k*3
//        println(listOf(cnt, n2,n3,n4))

        cnt += n2/5
        println(cnt)
    }
}
