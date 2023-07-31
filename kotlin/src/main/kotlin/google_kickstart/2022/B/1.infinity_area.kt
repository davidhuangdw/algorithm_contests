package google_kickstart.`2022`.B
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, A, B) = input_longs()
        var sum = .0
        var x = R
        while(x > 0){
            sum += x * x
            x *= A
            sum += x * x
            x /= B
        }
        sum *= Math.PI
        println("Case #${it}: $sum")
    }
}
