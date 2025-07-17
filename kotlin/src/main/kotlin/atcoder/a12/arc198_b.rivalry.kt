package atcoder.a12

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
    val (T) = inputInts()
    repeat(T) {
        val (x, y, z) = inputInts()
        println(if( z <= x && x*2 >= y
            && (y%2==0 || z>0)) "Yes" else "No") // bug: y must even
    }
}
