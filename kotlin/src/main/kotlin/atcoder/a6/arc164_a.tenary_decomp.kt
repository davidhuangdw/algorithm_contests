package atcoder.a6

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
    fun inputLongs() = readLine()!!.split(' ').map(String::toLong)

    repeat(inputInts()[0]){
        val (N, K) = inputLongs()
        var cnt = 0L
        var r = N
        while(r > 0){
            cnt += r % 3
            r/=3
        }
        println(if(cnt <= K && (K-cnt) % 2 == 0L) "Yes" else "No")
    }
}
